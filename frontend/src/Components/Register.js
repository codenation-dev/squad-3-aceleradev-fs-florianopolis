import React,{Fragment} from 'react';
import '../static/register.css';

const Register = () => {
    return(<Fragment>
        <div class="row">
        <div class="col-12 col-md-6">
        <h2>Registro de email</h2>
        <form>
            <label for="Nome">Nome</label>
            <input type="text" id="Nome"/>
            <label for="email">E-mail</label>
            <input type="email" id="Email"/>
            <label for="Senha">Senha</label>
            <input type="password" id="Senha"/>
            <button type="submit">Register</button>
        </form>
        </div>
        <div class="col-12 col-md-6" id="Registred">
        <h2>Emails Ja Registrados</h2> 
        <table>
          <tr>
          <th>Nome</th>
          <th>Email</th>
          <th>Ações</th> 
          </tr> 
        </table>
        </div>
        </div>
    </Fragment>)
}

export default Register;