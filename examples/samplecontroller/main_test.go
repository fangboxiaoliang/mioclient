package main

import (
	"fmt"
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hiboot/pkg/app/web"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/starter/kube"
	samplecontrollerv1alpha1 "github.com/hidevopsio/mioclient/examples/samplecontroller/pkg/apis/samplecontroller/v1alpha1"
	clientset "github.com/hidevopsio/mioclient/examples/samplecontroller/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)


type sampleController struct {
	restConfig *kube.RestConfig
}

func newSampleController(restConfig *kube.RestConfig) *sampleController {
	return &sampleController{
		restConfig: restConfig,
	}
}

func (c *sampleController) Register() error {

	//clientSet, err := apiextensionsclient.NewForConfig(c.restConfig.Config)
	//	//if err != nil {
	//	//	log.Debugf("Error building example clientset: %s", err.Error())
	//	//	return err
	//	//}
	//
	//apiExt := clientSet.ApiextensionsV1beta1()
	//crd := apiExt.CustomResourceDefinitions()
	//list, err := crd.List(metav1.ListOptions{})
	//if err != nil {
	//	return err
	//}
	//log.Debugf("%v", list)

	//
	//crd := &apiextensionsv1beta1.CustomResourceDefinition{
	//	ObjectMeta: metav1.ObjectMeta{Name: "foos.samplecontroller.k8s.io"},
	//	Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
	//		Group:   "samplecontroller.k8s.io",
	//		Version: "v1alpha1",
	//		Scope:   apiextensionsv1beta1.NamespaceScoped,
	//		Names:   apiextensionsv1beta1.CustomResourceDefinitionNames{
	//			Plural: "foos",
	//			Kind:   "Foo",
	//		},
	//	},
	//}
	//
	//if _,err := clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd);err != nil {
	//	log.Debugf("Error building example clientset: %s", err)
	//	//return err
	//}

	return nil
}


func (c *sampleController) CreateCrd()error{
	//create foo crd
	exampleClient, err := clientset.NewForConfig(c.restConfig.Config)
	if err != nil {
		log.Debugf("Error building example clientset: %s", err)
		return err
	}
	foo := exampleClient.SamplecontrollerV1alpha1().Foos("default")
	fooName := "example-foo"

	// create foo
	replica := int32(1)
	f := &samplecontrollerv1alpha1.Foo{
		ObjectMeta: metav1.ObjectMeta{Name: "example-foo"},
		Spec:samplecontrollerv1alpha1.FooSpec{
			DeploymentName:	fooName,
			Replicas: &replica,
		},
	}

	if _,err:= foo.Create(f);err != nil {
		log.Debugf("Error : %v", err)
		//return err
	}

	// get one
	if _,err := foo.Get("example-foo",metav1.GetOptions{});err != nil {
		log.Debugf("Error : %v", err)
		return err
	}

	// list
	fooList, err := foo.List(metav1.ListOptions{})
	if err != nil {
		log.Debugf("Error : %v", err)
		return err
	}


	if len(fooList.Items) == 0 {
		err := fmt.Errorf("list foo faile.")
		log.Debugf("Error : %v", err)
		return err
	}

	//update
	replica = int32(2)
	f.Spec.Replicas = &replica
	if _,err := foo.Update(f);err != nil {
		log.Debugf("Error : %v", err)
		return err
	}

	//watch
	watchFoo,err := foo.Watch(metav1.ListOptions{})
	if err != nil {
		log.Debugf("Error : %v", err)
		return err
	}

	if _,ok := <-watchFoo.ResultChan();!ok {
		err := fmt.Errorf("foo watch faile")
		log.Debugf("Error : %v", err)
		return err
	}

	// delete
	if err := foo.Delete(fooName,&metav1.DeleteOptions{});err != nil {
		log.Debugf("Error : %v", err)
		return err
	}
	return nil
}

func init() {
	app.Component(newSampleController)
	log.SetLevel(log.DebugLevel)
}

func TestCrd(t *testing.T) {
	testApp := web.NewTestApplication(t).(app.ApplicationContext)

	sampleController := testApp.GetInstance("sampleController").(*sampleController)

	t.Run("|should register crd", func(t *testing.T) {
		sampleController.Register()
	})

	t.Run("|should create crd", func(t *testing.T) {
		sampleController.CreateCrd()
	})
}
