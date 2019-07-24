import React,{useEffect} from 'react';
import '../static/login.css';
import {connect} from 'react-redux';
import {actionLogin} from '../redux/actions';
import PropTypes from "prop-types";
import {withRouter} from 'react-router-dom';

const Login = (props) => {
   
    useEffect(() => {
        if(props.Logged){
            props.history.push('/dashboard')
        }
    })
   
    return (
    <div className="login col-12 col-md-4">
    <form onSubmit={e => e.preventDefault()}>
    <label htmlFor="Email">Email</label>
    <input required type="Email" id="Email"/>
    <label htmlFor="Password">Password</label>
    <input required type="password" id="Password"/>
    <button onClick={props.logIn} type="submit">Login</button>
    </form>
    </div>);
}

Login.propTypes = {
    logIn: PropTypes.func
}

const mapStateToProps = state => {
    return {
        Logged: state.value
    }
}

const mapDispatchToProps = dispatch => {
    return {
        logIn: () => dispatch({type:'LOGIN'})
    }
}

export default connect(mapStateToProps,mapDispatchToProps)(withRouter(Login));