terraform {
  # backend "remote" {
  #   organization = "BelleChiu"

  #   workspaces {
  #     name = "terra-house-1"
  #   }
  # }

  cloud {
    organization = "BelleChiu"

    workspaces {
      name = "terraform-cloud"
    }
  }

}

module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
  index_html_filepath = var.index_html_filepath
  error_html_filepath = var.error_html_filepath
}
