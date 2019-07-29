import React,{useEffect,useState} from "react"
import {connect} from "react-redux"
import { REQUEST_FETCH } from "../redux/actions";
import Loading from '../Helpers/Loading';

const Warns = (props) => {
    
    let [date,setDate] = useState(null)
    let [ctr,setCtr] = useState(false)

    useEffect(()=> {
        if (ctr===false){
            setCtr(true)
            props.GetWarns(date)
        }
    })
    
    let RenderMails = () => (props.Warn.emails?props.Warn.emails.map(x=>(<tr><td>{x.emailusuario}</td><td>{x.data}</td></tr>)):"")

    let RenderMe = () => (props.Warn?<div>
            <div>
               <h3>Alerta <span>#{props.Warn.id}</span></h3> 
               <h6>Data: {props.Warn.data}</h6>
               <h4>Seus Clientes que viraram Funcionarios publicos</h4>
                <ul>{props.Warn.lista.ClientesDoBanco||"Nenhum encontrado"}</ul>{/*.map(x=><li>{x}</li>) */}
               <h4>Pessoas que poderiam virar seus clientes!</h4>
               <ul>{props.Warn.lista.TopFuncionariosPublicos||"Nenhum encontrado"}</ul>
            </div>
                <h5>Emails Enviados</h5>
            <table>
                <tr>
                    <th>Usuario</th>
                    <th>Data do envio</th>
                </tr>
                {RenderMails()}
            </table>
        </div>:"")
    
    return (<Loading Loaded={RenderMe}/>)
}

const mapStateToProps = state => {
    return {
        Warn: state.API.Warn 
    }
}

const mapDispatchToProps = dispatch => {
    return {
        GetWarns: (date) => dispatch({type:REQUEST_FETCH,endpoint:"Warns",data:date})
    }
}

export default connect(mapStateToProps,mapDispatchToProps)(Warns)