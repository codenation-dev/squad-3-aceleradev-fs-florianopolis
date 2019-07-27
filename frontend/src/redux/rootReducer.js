const Logged = (state = {value:false,firstLogin:false}, action) => {
    switch (action.type) {
        case 'LOGOUT':
           return {value:false,firstLogin:false};
        case 'LOGIN':
            return {value:true,firstLogin:true};
        case 'LOGGED':
            return {value:true,firstLogin:false}
        default:
            return state
    }
}

export default Logged;