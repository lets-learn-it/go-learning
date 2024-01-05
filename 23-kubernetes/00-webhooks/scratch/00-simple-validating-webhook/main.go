package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ServeValidatePods(w http.ResponseWriter, r *http.Request) {
	log.Println("Got validation request.")
	decoder := json.NewDecoder(r.Body)

	var t admissionv1.AdmissionReview

	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}

	log.Println("Kind: ", t.Request.Kind.Kind)

	if t.Request.Kind.Kind != "Pod" {
		log.Println("Not a pod. ignoring.")
	}

	var p corev1.Pod
	if err := json.Unmarshal(t.Request.Object.Raw, &p); err != nil {
		log.Println("failed to unmarshal pod.", err)
	}

	res := &admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AdmissionReview",
			APIVersion: "admission.k8s.io/v1",
		},
		Response: &admissionv1.AdmissionResponse{
			UID:     t.Request.UID,
			Allowed: true,
			Result: &metav1.Status{
				Status: "Success",
			},
		},
	}
	if p.Labels["owner"] != "PSKP-95" {
		res.Response.Allowed = false
		res.Response.Result.Status = "Failure"
		res.Response.Result.Message = "label owner with value PSKP-95 not exists."

		log.Println("owner label is absent.")
	}

	w.Header().Set("Content-Type", "application/json")
	rout, err := json.Marshal(res)
	if err != nil {
		log.Println("Error while marshelling")
	}

	fmt.Fprintf(w, "%s", rout)
}

func ServeHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	http.HandleFunc("/validate-pods", ServeValidatePods)
	http.HandleFunc("/health", ServeHealth)

	log.Println("Starting web server...")
	fmt.Println(http.ListenAndServeTLS(":8080", "./manifests/ca.crt", "./manifests/ca.key", nil))
}
