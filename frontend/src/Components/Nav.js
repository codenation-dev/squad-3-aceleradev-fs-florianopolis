import React, { useState, useEffect } from "react";
import "../static/nav.css";
import { Link } from "react-router-dom";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import actionLogout from "../redux/actions";

const Nav = props => {
  
  const [expander, setExpander] = useState(false);
  const [loginAnim, setAnim] = useState(false);

  const handleClick = () => {
     setExpander(!expander);}

  const Collapser = (value) => (value?"Expand":"")
  
  const FirstLoginAnimation = (normal,value) => (value?`${normal} FirstLoginAnimation`:normal)

  useEffect(()=>{
	  if(props.isLogged.firstLogin){
		  setAnim(true)
		  props.LogConfirm()
	  }
  })

	const render = props => {
		return !props.isLogged.value ? (
			<nav>
				<Link to="/" className="Brand">
					<h1>Uati</h1>
				</Link>
			</nav>
		) : (
			<nav className="Logged">
				<Link to="#" className={FirstLoginAnimation("Brand",loginAnim)}>
					<h1>Uati</h1>
				</Link>
				<div className={FirstLoginAnimation("navLinks",loginAnim)}>
					<span onClick={handleClick} className="Expander"><Link to="#" >Opções</Link></span>
          <ul className={Collapser(expander)}>
            <Link onClick={handleClick} to="/dashboard">
						<li>Dashboard</li>
            </Link >
			<Link onClick={handleClick} to="/mailregister">
				<li>Email</li>
			</Link>
			<Link onClick={handleClick} to="/warn">
						<li>Alertas</li>
			</Link>
			<Link onClick={handleClick} to="#" >
			<li onClick={props.Logout}>Logout</li>
			</Link>
					</ul>
				</div>
			</nav>
		);
	};
  
  return render(props);
};

Nav.propTypes = {
	Logout: PropTypes.func,
	isLogged: PropTypes.bool,
	history: PropTypes.object
};

const mapDispatchToProps = dispatch => {
	return { Logout: () => dispatch({ type: "REQUEST_LOGOUT" }),
			 LogConfirm: () => dispatch({type:"REFRESH_LOGIN"}) };
};

const mapStateToProps = state => {
	const Logged = state.Login;
	return {
		isLogged: Logged
	};
};

export default connect(
	mapStateToProps,
	mapDispatchToProps
)(Nav);
