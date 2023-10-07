terraform {
  required_providers {
    terratowns ={
      source = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
  # backend "remote" {
  #   organization = "BelleChiu"

  #   workspaces {
  #     name = "terra-house-1"
  #   }
  # }

  # cloud {
  #   organization = "BelleChiu"

  #   workspaces {
  #     name = "terra-house-1"
  #   }
  # }

}

provider "terratowns" {
  endpoint = "http://localhost:4567/api"
  user_uuid="f39cf958-47a2-40c7-96cb-e8479413d33d" 
  token="9b49b3fb-b8e9-483c-b703-97ba88eef8e0"
}

# module "terrahouse_aws" {
#   source = "./modules/terrahouse_aws"
#   user_uuid = var.user_uuid
#   bucket_name = var.bucket_name
#   index_html_filepath = var.index_html_filepath
#   error_html_filepath = var.error_html_filepath
#   assets_path = var.assets_path
#   content_version = var.content_version
# }

resource "terratowns_home" "home"{
  name = "How to cook and bake!!"
  description = <<DESCRIPTION
  Cooking and baking are culinary techniques used to prepare a wide variety of foods.
  Remember that cooking and baking can be both simple and complex, and 
  there's a wide range of dishes you can prepare. As you gain experience, 
  you'll become more comfortable with different techniques and flavors.
DESCRIPTION
#domain_name = module.terratowns_aws.cloudfront_url
  domain_name = "af5511d.cloudfront.net"
  town = "cooker-cove"
  content_version = 1
}