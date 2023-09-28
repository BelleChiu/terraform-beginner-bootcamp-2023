# Terraform Beginner Bootcamp 2023 - Week 1

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





