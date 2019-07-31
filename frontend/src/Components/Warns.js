import React,{Fragment,useEffect,useState} from "react"
import {connect} from "react-redux"
import { REQUEST_FETCH } from "../redux/actions";
import Loading from '../Helpers/Loading';
import '../static/warn.css'

const Filter = (props) => {
    const [id,setId] = useState(1)
    let handleChange = e => {
        setId(e.target.value)
    }
    return(
    <div className="FilterWarn">    
    <h1>Filtrar Alerta</h1>
    <label for="alertaID" >ID do alerta</label>
    <input id="alertaID" value={id} type="number" onChange={handleChange} min="1" max="6"/>
    <button onClick={()=>{props.handle(id)}} >Pesquisar</button>
    </div>
)
}

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

    let toList = (List) => {
        if (List) {
            return List.map(x=><li>{x}</li>)
        } else {
            return ("Nenhum encontrado")
        }
    }

    let AllOrNone = () => (props.Warn.id>0?<div id="WarnG">
    <div>
       <h3><span>Alerta</span> #{props.Warn.id}</h3> 
       <h4><span>Data: </span>{props.Warn.data}</h4>
       <h4><span>Seus Clientes que viraram Funcionarios publicos</span></h4>
        <ul>{toList(props.Warn.lista.ClientesDoBanco)}</ul>
       <h4>Pessoas que poderiam virar seus clientes!</h4>
       <ul>{toList(props.Warn.lista.TopFuncionariosPublicos)}</ul>
    </div>
        <h5><span>Emails Enviados</span></h5>
    <table>
        <tr>
            <th>Usuario</th>
            <th>Data do envio</th>
        </tr>
        {RenderMails()}
    </table>
</div>:<h1>Nenhum dado encontrado para o ID referido!</h1>)

    let RenderMe = () => (props.Warn?<div className="GeneralDiv"><Filter handle={(x)=>props.GetByID({id:parseInt(x)})}/>
    {AllOrNone()}
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
        GetWarns: (date) => dispatch({type:REQUEST_FETCH,endpoint:"Warns",data:date}),
        GetByID: (IDs) => dispatch({type:REQUEST_FETCH,endpoint:"Warns",args:IDs})
    }
}

export default connect(mapStateToProps,mapDispatchToProps)(Warns)