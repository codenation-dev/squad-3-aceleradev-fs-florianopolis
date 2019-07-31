import React,{useEffect,useState} from 'react';
import '../static/login.css';
import {connect} from 'react-redux';
import {actionLogin} from '../redux/actions';
import PropTypes from "prop-types";
import {withRouter} from 'react-router-dom';

const Login = (props) => {
    
    let [usermail,setUser] = useState("")
    let [password,setPass] = useState("")

    useEffect(() => {
        if(props.Logged){
            props.history.push('/dashboard')
        }
    })
   
    const handleChange = e => {
        switch (e.target.id) {
            case "Email":
                setUser(e.target.value)
                break;
            case "Password":
                setPass(e.target.value)
                break;
            default:
                break;
        }
    }

    let FormFields = {usermail:usermail,password:password}

    return (
    <div className="login col-12 col-md-4">
    <form onSubmit={e => e.preventDefault()}>
    <label htmlFor="Email">Email</label>
    <input onChange={handleChange} required type="Email" id="Email"/>
    <label htmlFor="Password">Password</label>
    <input onChange={handleChange} required type="password" id="Password"/>
    <button onClick={() => props.logIn(FormFields)} type="submit">Login</button>
    </form>
    </div>);
}

Login.propTypes = {
    logIn: PropTypes.func
}

const mapStateToProps = state => {
    return {
        Logged: state.Login.value,
        Loading: state.API.Loading
    }
}

const mapDispatchToProps = dispatch => {
    return {
        logIn: (FormFields) => dispatch({type:'REQUEST_LOGIN',loginFields: FormFields})
    }
}


export default connect(mapStateToProps,mapDispatchToProps)(withRouter(Login));