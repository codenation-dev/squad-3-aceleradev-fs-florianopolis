import React, { useEffect } from "react";
import { connect } from "react-redux";
import actionLogout from "../redux/actions";

const UnauthRedirect = (props) => {
   return (props.Logged?props.Auth:props.Unath)
}

const mapStateToProps = state => {
    let Logged = state.Login.value;
    return {Logged:Logged}
}

export default connect(mapStateToProps)(UnauthRedirect)