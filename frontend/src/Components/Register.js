import React, { useEffect, useState, Fragment, Link } from "react";
import "../static/register.css";
import { connect } from "react-redux";
import { REQUEST_FETCH } from "../redux/actions";
import Loading from "../Helpers/Loading";
import Edit from "./Edit";
import { EditIcon, Trash } from "../Helpers/Icons";
import Remove from './Remove'

const Register = props => {
	let [name, setName] = useState("");
	let [email, setEmail] = useState("");
	let [ctr, setCtr] = useState(false);

	useEffect(() => {
		if (ctr === false) {
			setCtr(true);
			props.userMailList();
		}
	});

	const handleChange = e => {
		switch (e.target.id) {
			case "Nome":
				setName(e.target.value);
				break;
			case "Email":
				setEmail(e.target.value);
				break;
			default:
				break;
		}
	};

	const handleClick = x => {
		props.userSelect(x);
	};
	//<Route path="/mailedit" exact component={EditRoute}/>
	let RenderMails = () =>
		props.Users
			? props.Users.map(x => (
					<tr>
						<td>{x.Name}</td>
						<td>{x.Mail}</td>
						<td onClick={() => handleClick(x)}>
								<span onClick={() => props.userOptions("EditModal")}>
									{EditIcon()}
								</span>
								<span onClick={() => props.userOptions("RemoveModal")}>
									{Trash()}
                                </span> 
						</td>
					</tr>
			  ))
			: "";

	let Fields = { Name: name, Mail: email };
	const RenderMe = () => (
		<Fragment>
			<Edit />
            <Remove />
			<div class="row">
				<div class="col">
					<h2 class="boxTitle">Registro de email</h2>
					<form>
						<label for="Nome">Nome</label>
						<input onChange={handleChange} required type="text" id="Nome" />
						<label for="email">E-mail</label>
						<input onChange={handleChange} required type="email" id="Email" />
						<button
							type="submit"
							onClick={e => {
								e.preventDefault();
								props.userRegister(Fields);
							}}
						>
							Register
						</button>
					</form>
				</div>
				<div class="col" id="Registred">
					<h2 class="boxTitle">Emails Ja Registrados</h2>
					<table>
						<thead>
							<th>Nome</th>
							<th>Email</th>
							<th>Ações</th>
						</thead>
						<tbody>{RenderMails()}</tbody> 
					</table>
				</div>
			</div>
		</Fragment>
	);
	return <Loading Loaded={RenderMe} />;
};

const mapStateToProps = state => {
	return {
		Users: state.API.UsermailList,
		Loading: state.API.Loading
	};
};

const mapDispatchToProps = dispatch => {
	return {
		userMailList: () => dispatch({ type: REQUEST_FETCH, endpoint: "Users" }),
		userRegister: Fields =>
			dispatch({ type: REQUEST_FETCH, endpoint: "UserAdd", args: Fields }),
		userSelect: x => dispatch({ type: "LOAD_USER", data: x }),
		userOptions: x => dispatch({ type: "USER_MODAL", modal: x })
	};
};

export default connect(
	mapStateToProps,
	mapDispatchToProps
)(Register);
