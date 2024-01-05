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

func ServeMutatePods(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	var p corev1.Pod
	if err := json.Unmarshal(t.Request.Object.Raw, &p); err != nil {
		log.Println("failed to unmarshal pod.", err)
	}

	// considering there is no en
	patch := ""
	if p.Spec.Containers[0].Env == nil {
		patch = `[{"op":"add","path":"/metadata/labels/myLabel","value":"demo"},{"op":"add","path":"/spec/containers/0/env","value":[{"name":"envvar","value":"PPPP"}]}]`
	} else {
		patch = `[{"op":"add","path":"/metadata/labels/myLabel","value":"demo"},{"op":"add","path":"/spec/containers/0/env/0","value":{"name":"envvar","value":"PPPP"}}]`
	}

	patchType := admissionv1.PatchTypeJSONPatch
	res := &admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AdmissionReview",
			APIVersion: "admission.k8s.io/v1",
		},
		Response: &admissionv1.AdmissionResponse{
			UID:       t.Request.UID,
			Allowed:   true,
			PatchType: &patchType,
			Patch:     []byte(patch),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	rout, err := json.Marshal(res)
	if err != nil {
		log.Println("Error while marshelling")
	}
	log.Println(string(rout))
	fmt.Fprintf(w, "%s", rout)
}

func ServeHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	http.HandleFunc("/mutate-pods", ServeMutatePods)
	http.HandleFunc("/health", ServeHealth)

	log.Println("Starting web server...")
	fmt.Println(http.ListenAndServeTLS(":8080", "./manifests/ca.crt", "./manifests/ca.key", nil))
}
