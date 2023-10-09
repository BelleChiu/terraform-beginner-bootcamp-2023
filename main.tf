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

  cloud {
    organization = "BelleChiu"

    workspaces {
      name = "terra-house-1"
    }
  }

}

provider "terratowns" {
  endpoint = var.terratowns_endpoint
  user_uuid= var.teacherseat_user_uuid
  token=var.terratowns_access_token

}

module "home_hometown_hosting" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  public_path = var.hometown.public_path
  content_version = var.hometown.content_version
}

resource "terratowns_home" "home"{
  name = "My HomeTown!!"
  description = <<DESCRIPTION
  Tainan City is a special municipality in southern Taiwan facing the Taiwan Strait on its western coast. 
  Tainan is the oldest city on the island and also commonly known as the "Capital City"[II] for its over 200 years of history
  as the capital of Taiwan under Koxinga and later Qing rule
DESCRIPTION
  domain_name = module.home_hometown_hosting.domain_name
  #domain_name = "af5511d.cloudfront.net"
  town = "missingo"
  content_version = var.hometown.content_version
}

module "home_cookandbake_hosting" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  public_path = var.cookandbake.public_path
  content_version = var.cookandbake.content_version
}

resource "terratowns_home" "home_cookandbake"{
  name = "How to cook and bake"
  description = <<DESCRIPTION
  Discover Tainan's vibrant street food culture and traditional dishes. 
  Learn to cook local recipes and savor the unique flavors of this city. 
  Moreover, indulge in Tainan's sweet treats and pastries. Whether you're 
  a beginner or a baking enthusiast, you'll find delicious recipes and tips to satisfy your cravings.
DESCRIPTION
  domain_name = module.home_cookandbake_hosting.domain_name
  #domain_name = "af5511d.cloudfront.net"
  town = "cooker-cove"
  content_version = var.cookandbake.content_version
}