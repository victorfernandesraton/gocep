const getCep = async (cep) => {
    const req = await fetch(`http://localhost:8080/cep/${cep}`)
    console.log("Status Code" + req.status)
    const data = await req.json()
    return data
}

const data = await getCep('01001000')

console.log(data)