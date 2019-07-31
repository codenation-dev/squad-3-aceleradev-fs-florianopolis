import axios from "axios";

const BASE_URL = "http://localhost:8921";

export const login = async (LoginFields) => {
    let URL = (BASE_URL+"/auth")
    let tryIt = await axios.post(`${URL}`,{
        usermail: LoginFields.usermail,
        password: LoginFields.password
    })
    console.log(tryIt.data)
    return tryIt.data
}

export const request = async (token,endpoint,...args) => {
    let Addr = "";
    switch (endpoint) {
        case 'Users':
            Addr = "/mails"
            break;
        case 'Warns':
            Addr = "/warn"
            break;
        case 'UserAdd':
            Addr = "/mails/add"
            break;
        case 'Update':
            Addr = `/mails/${args[0].id}/update`
            break;
        case 'Delete':
            Addr = `/mails/${args[0].id}/delete`
            break;
        case 'DataScience':
            Addr = '/tables';
            break;
        default:
            break;
    }
    let tokenHeader = {"Access-Token":token}
    let URL = (BASE_URL+Addr);
    console.log(URL)
    console.log({...args[0]})
    let tryIt = await axios.post(`${URL}`,{...args[0]},{headers:tokenHeader})
    return tryIt.data
}

export const uploadfile = async (token,list) => {
    let tokenHeader = {"Access-Token":token}
    let URL = (BASE_URL+"/upload");
    console.log(list)
    let tryIt = await axios.post(`${URL}`,list,{headers:tokenHeader})
    return tryIt.data
}