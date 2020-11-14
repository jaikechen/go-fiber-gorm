import Axios from "axios"
const url = 'http://localhost:3000/api/book'
async function get(id) {
    await execute(async()=>{
        const result = await Axios.get(url+'/' + id)
        console.log(result.data)
    })

}
async function post(item: any) {
    await execute(async () => {
        const result = await Axios.post(url, item)
        console.log(result.data)
    })
}

async function execute(cb: any) {
    try {
        await cb()
    }
    catch (error) {
        if (error.response) {
            console.log(error.response.data)
        } else {
            console.log(error)
        }
    }
}

get('')
//post({ title: '1984', author:'Owel',rating:'100' })
