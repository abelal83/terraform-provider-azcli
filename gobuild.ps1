$env:GOOS = "windows"
$env:GOARCH = "amd64"
remove-item ".\terraform-provider-azcli_v0.0.6_x64.exe" -Force -ErrorAction SilentlyContinue
go build -o "terraform-provider-azcli_v0.0.6_x64.exe"
#Remove-Item .\.terraform -Force -Recurse
#remove-item .\terraform.tfstate -Force
#terraform init
#terraform.exe plan 

$env:GOOS = "linux"
remove-item ".\terraform-provider-azcli_v0.0.6_x64" -Force -ErrorAction SilentlyContinue
go build -o "terraform-provider-azcli_v0.0.6_x64"
