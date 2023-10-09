output "bucket_name" {
  description = "Bucket name for our static website hosting"
  value = module.home_hometown_hosting.bucket_name
}

output "s3_website_endpoint" {
  description = "S3 Static Website Hosting Endpoint"
  value=module.home_hometown_hosting.website_endpoint
}


output "domain_name" {
  value = module.home_hometown_hosting.domain_name
}