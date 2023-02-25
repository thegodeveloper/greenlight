for ($i=1; $i -le 6; $i++) {
    Invoke-RestMethod -Uri http://localhost:4000/v1/healthcheck
}
