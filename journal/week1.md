# Terraform Beginner Bootcamp 2023 - Week 1

## Fixing Tags

Local delete a tag
```sh
git tag -d <tag_name>
```

Remotely delete tag
```sh
 git tag -d v2.0
 ```

Checkout the commit that you want to retag. Grab the sha from your Github history.

```sh
git checkout <SHA>
git tag M.M.P
git push --tags
git checkout main
```



[How to Delete Local and Remote Tagas on Git](https://devconnected.com/how-to-delete-local-and-remote-tags-on-git/)



## Root Module Structure

Our root module structure is as follows:

```
PROJECT_ROOT
│
├── variables.tf      # stores the structure of input variables
├── main.tf           # everything else.
├── providers.tf      # defined required providers and their configuration
├── outputs.tf        # stores our outputs
├── terraform.tfvars  # the data of variables we want to load into our Terraform project
└── README.md         # required for root modules
```

[Standard Module Structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

## Terraform and Input Variables
### Terraform Cloud Variables

In terraform we can set two kind of variables:
- Environment Variables - those that you would set in your bash terminal eg. AWS credentials
- Terraform Variables - those that you can would normally set in your tfvars file

We can set Terraform Cloud Variable to be sensitive so they are not shown visibly in the UI.

### Loading Terraform Input Variables

[Terraform Input Variables](https://developer.hashicorp.com/terraform/language/values/variables)

### var flag
We can user the `-var` flag to set an input variable or override a variable in the tfvars file eg.`terraform -var user_ud="my-user_id`

Or, when running terraform plan and terraform apply, and use multiple `-var` as follow:
 
 We can create main.tf with the following content
 ``` 
 terraform {
    .....
 }

 provider "aws"{
    .....
 }

 variable "instance_type" {
  description = "EC2 instance type"
  type        = string
}

variable "cidr_block" {
  description = "CIDR block for VPC"
  type        = string
}

resource "aws_vpc" "example" {
  cidr_block = var.cidr_block
}

resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = var.instance_type
}

 ```
 Then, we execute the following bash script:
 ```bash
  terraform plan -var "instance_type=t3.micro" -var "cidr_block=10.1.0.0/16"
 ```


### var -file flag

We can set a lots of variables in a variable definitions file (with a filename ending in either .tfvars or .tfvars.json) and then execute the following command line with the file.

```terraform apply -var-file="ex_var.tfvars"```


### terraform.tvfars

This is the default file to load in terraform variables in blunk 

### auto.tfvars

In Terraform, .auto.tfvars is a special file that can be used to automatically load variable values when running terraform apply or terraform plan. This file allows you to specify default values for variables without the need to pass them on the command line using the -var flag. The .auto.tfvars file should be located in the same directory as your Terraform configuration files.

e.g.
- Create a Terraform configuration file, for example, main.tf, with variables:
``` hcl
variable "instance_type" {
  description = "EC2 instance type"
  type        = string
}

variable "rn" {
  description = "AWS region"
  type        = string
}
```
- Create .auto.tfvars file in the same directory
```
instance_type = "t2.micro"
rn = "us-west-2"
```
- run `terraform plan` or `terraform apply`, Terraform will automatically load variable values from the `.auto.tfvars`

### order of terraform variables

Terraform loads variables from multiple sources in a specific order, with later sources potentially overwriting values from earlier sources. The order of variable loading is as follows:

Terraform loads variables in the following order, with later sources taking precedence over earlier ones:

- Environment variables. For example, if you have a variable named **my_variable**, Terraform will look for an environment variable **named TF_VAR_my_variable** and use its value.

- The terraform **.tfvars** file, if present.

- The terraform **.tfvars.json** file, if present.

- Any ***.auto.tfvars** or ***.auto.tfvars.json** files, processed in lexical order of their filenames.

- Any `-var` and `-var-file` options on the command line, in the order they are provided. (This includes variables set by a Terraform Cloud workspace.)



##   Dealing with Configuration Drift

## What happens f we lose our state file?
If you lose your statefile, you most likey have to tear down all your cloud infrastructure manually.

You can use `terraform import` but it won't for all cloud resources. You need to check the terraform providers documentation for which resources support import.

## Fix Missing Resource with Terraform import

`terraform import aws_s3_bucket.bucket bucket-name`

[Terraform Import](https://developer.hashicorp.com/terraform/cli/import)
[AWS S3 Bucket Import](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#import)

### Fix Manual Configuration 
If someone goes and delete or modifies cloud resource manually through ClickOps

If we run `terraform plan` is with attempt to put our infrastructure baCK into the expected state fixing Configuration Drift.
### Fix using Terraform Refresh

```sh
terraform apply -refresh-only -auto-approve
```

## Terraform Modules

### Terraform Module Structure

It is recommend to place modules in a `modules` directory when locally developing modules but you can name it whatever you want.

```
PROJECT_ROOT
|
├── variables.tf      
├── main.tf             
├── outputs.tf        
├── README.md
└──  modules/
    └── nestedA/
          ├── README.md    
          ├── variables.tf
          ├── main.tf
          ├── LICENSE
          └── outputs.tf
```

## Passing Input Variables
We can pass input variables to our module.

The module has to declare the terraform variables in its own variables.tf

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
}

```

### Modules Sources

Using the source we can import the module from various palces eg:
- locally
- Github
- Terraform Registry

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
}

```

[Modules Sources](https://developer.hashicorp.com/terraform/language/modules/sources)

## Considerations when using GhatGPT to write Terraform 

LLMs such as ChatGPT may not be trained on the latest documentation or information about Terraform.

It may likely produce older examples that cloud be deprecated. Often affecting providers.

eg. You can write "what is the lastest aws terraform provider version" in prompt.   
And ChapGpt give the follwoing content:
As of my last knowledge update in **September 2021**, I do not have access to real-time data. To find **the latest version of the AWS Terraform provider**, I recommend checking the official **Terraform Registry** or the **HashiCorp website**.

## Working with Files in Terraform

### Can and File function

There are built in terraform function to check the exitance of file.

```tf
condition     = can(file(var.index_html_filepath))
```

- [can function](https://developer.hashicorp.com/terraform/language/functions/can)
- [file function](https://developer.hashicorp.com/terraform/language/functions/file)

### Filemd5

[filemd5](https://developer.hashicorp.com/terraform/language/functions/filemd5)

### Path Variable

In terraform there is special variable called `path` that allows us to reference local paths :
- path.module = get the path for the current module
- path.root = get the path for the root module

To check where the expression is placed. we can use `terraform console`

[Special Path Variable]https://developer.hashicorp.com/terraform/language/expressions/references#path-module)

resource "aws_s3_object" "index_html" {
  bucket = aws_s3_bucket.website_bucket.bucket
  key    = "index.html"
  source = "${path.root}/public/index.html"
}

### Terraform locals

Locals allows us to define local variables.
It can be very useful when we need to terraform data into another format and have referenced a variables.



```tf
locals {
    s3_origin_id = "MyS3Origin"
}
```

[Local Values](https://developer.hashicorp.com/terraform/language/values/locals)
### Terraform DataSource

This allows to use source data from cloud resources.

This is useful when we want to reference cloud resources without importing them.

```tf
data "aws_caller_identity" "current" {}

output "account_id" {
  value = data.aws_caller_identity.current.account_id
}
```
[Data Sources:aws_caller_identify](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity)

[Data Sources](https://developer.hashicorp.com/terraform/language/data-sources)

## Working with JSON

We use the jsonencode to create the json policy inline in the hcl.

```tf
> jsonencode({"hello"="world"})
{"hello":"world"}

```

[jsonencode](https://developer.hashicorp.com/terraform/language/functions/jsonencode)


### Changing the Lifecycle of resources 

[Meta Arguments Lifecycle](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle)

## Terraform Data

Plain data values such as **Local Values** and **Input Variables** don't have any side-effects to plan against and so they aren't valid in replace_triggered_by. You can use terraform_data's behavior of planning an action each time input changes to indirectly use a plain value to trigger replacement.

```tf
variable "version" {
  default = 1
}

resource "terraform_data" "replacement" {
  input = var.version
}

# This resource has no convenient attribute which forces replacement,
# but can now be replaced by any change to the revision variable value.
resource "example_database" "test" {
  lifecycle {
    replace_triggered_by = [terraform_data.replacement]
  }
}

```
[terraform_data](https://developer.hashicorp.com/terraform/language/resources/terraform-data)


## Provisioners 

Provisioners allow you to execute commands on compute instances eg. AWS CLI command

They are not recommended for use by Hashicorp because Configuration Management tools such as Ansible are better fit, but the functionality exists.

[Provisioners](https://developer.hashicorp.com/terraform/language/resources/provisioners/syntax)

### Local-exec

This will execute commands on the machine running the terraform commands eg. plan or apply

```tf
resource "aws_instance" "web" {
  # ...

  provisioner "local-exec" {
    command = "echo ${self.private_ip} >> private_ips.txt"
  }
}
```

https://developer.hashicorp.com/terraform/language/resources/provisioners/local-exec


### Remote-exec


This will execute commands n a machine wich your target. You need to provide credentials such as ssh to get into the machine.

```tf
resource "aws_instance" "web" {
  # ...

  # Establishes connection to be used by all
  # generic remote provisioners (i.e. file/remote-exec)
  connection {
    type     = "ssh"
    user     = "root"
    password = var.root_password
    host     = self.public_ip
  }

  provisioner "remote-exec" {
    inline = [
      "puppet apply",
      "consul join ${aws_instance.web.private_ip}",
    ]
  }
}

```
https://developer.hashicorp.com/terraform/language/resources/provisioners/remote-exec

## For Each Expressions 

For and For Each allows us to enumerate over complex data types

```tf
[for s in var.list : upper(s)]

```
```tf
resource "aws_iam_user" "the-accounts" {
  for_each = toset( ["Todd", "James", "Alice", "Dottie"] )
  name     = each.key
}
```

This is mostly useful when you are creating multiples of a cloud resource and you want to reduce the amount of repetitive terraform code.

[For Expressions](https://developer.hashicorp.com/terraform/language/expressions/for)

[For Each Expressions](https://developer.hashicorp.com/terraform/language/meta-arguments/for_each)
