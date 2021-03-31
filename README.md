# CRD: Custom Resource Definition
# hello there 
Steps:

1) follow [This](https://www.openshift.com/blog/kubernetes-deep-dive-code-generation-customresources)

2) create the folders as the structure given in the link. Here only the directories that need to be created are:
   - hack
     - boilerplate.go.txt
   - pkg>apis>stable.example.com>v1alpha1
        - doc.go
        - register.go
        - types.go
   - pkg>apis>stable.example.com
        - register.go
    
These are the folders and files that need to be created before code generation.
      
3) the hack> update-codegen.sh file needs to be made an executable.
```
 chmod +x hack/update-codegen.sh 
```
4) go mod init
   
5) go mod go mod tidy && go mod vendor
   
6) ```
   ./hack/update-codegen.sh
   ```
This will generate the files

7) Copy the controller-gen in the bin folder

8) make the controller-gen as executable file

9) run:
```
make
```
This will create the yaml file.

10) 
```
kubectl apply -f <generated yaml file>
```
This will make the resource known to kubernetes api.

