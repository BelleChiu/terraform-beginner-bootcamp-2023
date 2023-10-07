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
  endpoint = var.terratowns_endpoint
  user_uuid= var.teacherseat_user_uuid
  token=var.terratowns_access_token

}

module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.teacherseat_user_uuid
  index_html_filepath = var.index_html_filepath
  error_html_filepath = var.error_html_filepath
  assets_path = var.assets_path
  content_version = var.content_version
}

resource "terratowns_home" "home"{
  name = "My HomeTown!!"
  description = <<DESCRIPTION
  Tainan City is a special municipality in southern Taiwan facing the Taiwan Strait on its western coast. 
  Tainan is the oldest city on the island and also commonly known as the "Capital City"[II] for its over 200 years of history
  as the capital of Taiwan under Koxinga and later Qing rule
DESCRIPTION
  domain_name = module.terrahouse_aws.cloudfront_url
  #domain_name = "af5511d.cloudfront.net"
  town = "missingo"
  content_version = 1
}