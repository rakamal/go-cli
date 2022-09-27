package main

import (

	"bytes"
//	"strings"
	"fmt"
	"log"
	"os"
	"os/exec"
)
var vpcname, group, imageid, instancetype, cmd1, keyname = "vpc-9df038e0", "any-grp", "ami-f0e7d19a", "t2.micro", "sg-0a94428cbd8cc8bfa", "awskey1"
func createAWSVM() {
	//aws := "aws"
	//ec2 := "ec2"
	//createsecuritygroup := "create-security-group"
	//vpcname_ := "--vpc-name"
	//group_ := "--group-name"
	//description := "awsvm"
	//runinstance := "run-instances"
	//instancetype_ := "--instance-type"
	//imageid_ := "--image-id"
	//securitygroupid := "--security-group-id"
	//createkeypair := "create-key-pair"
	//keyname_ := "--key-name"
	//query := "--query"
	//output := "--output"
	//string1 := "'KeyMaterial'"
	//text := "text"
	fmt.Println("Creating directory")
	if err := os.MkdirAll("/home/Abi-dev/program/.md/vm", 0777); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created directory")
	fmt.Println("Creating key-pair values..")
	cmd := exec.Command("bash","-c","aws ec2 create-key-pair --key-name awskey1 --query 'KeyMaterial' --output text")
	log.Printf("%v",cmd)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		log.Printf("STDERR of install  command : %v\n", stderr.String())
	}
	log.Printf("Output of install command : %v\n", string(out))
	f, err := os.Create("/home/Abi-dev/program/.md/vm/llss.pem")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(string(out))

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("done")
	fmt.Println("Started creating sg")
	cmd1 := exec.Command("bash","-c","aws ec2 create-security-group --group-name isec --description 'mg' --vpc-id vpc-9df038e0")
	cmd1.Stderr = &stderr
	outp, err := cmd1.Output()
	if err != nil {
		log.Printf("STDERR of sg command : %v\n", stderr.String())
	}
	myvar := string(outp)
	anyVar := "aws ec2 run-instances --image-id ami-052efd3df9dad4825 --count 1 --instance-type t2.micro --key-name awskey1 --security-group-ids " + myvar
	log.Printf("Output of sg command : %s\n", myvar)
     fmt.Println("Created security grp")
     cmd2 := exec.Command("bash","-c","$anyVar")
     log.Printf("install_kubectl:CMD>>:%v\n", cmd2)
	cmd2.Stderr = &stderr
	outpu, err := cmd2.Output()
	if err != nil {
		log.Printf("STDERR of vm command : %v\n", stderr.String())
	}
	log.Printf("Output of vm install command : %v\n", string(outpu))

}
func deleteAWSVM() {

}
func main() {
	createAWSVM()
}
