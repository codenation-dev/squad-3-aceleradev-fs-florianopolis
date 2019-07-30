import React,{useState,useEffect, Fragment} from 'react';
import '../static/dashboard.css';
import {connect} from 'react-redux';
import { REQUEST_FETCH } from "../redux/actions";
import Charter from './Charts';
import Loading from '../Helpers/Loading'

const Selecter = (props) => {
    let selfHandler = () => {
        props.onClick(props.value)
    }
    let active = () => (props.active?"active":"")
return(<li className={active()} onClick={selfHandler}>{props.value}</li>)}

const SelecterList = (props) => {
    let itensList = props.Options.map(x=><Selecter active={props.Selected===x.value} value={x.value} onClick={props.onPress} />)
    let expander = () => (props.Expanded?"Expand":"")
    return (<ul className={expander()}>{itensList}</ul>)}

const ItemList = (props) => {
    let itensList = props.Options.map(x=><Item selected={props.Selected===x.value} Render={<Fragment><h2>{x.value}</h2><p className="desc"><span className="desc">Descrição :&nbsp;</span> {x.desc}</p><div className="Charter"><Charter type={x.type} data={x.dsmap} /></div></Fragment>}/>)
    return (<div className="FilterContainer">{itensList}</div>)
}    

const Item = (props) => {
    let show = () => (props.selected?"ActiveItem":"FilterItem")
    return(<div className={show()}>{props.Render}</div>)
}    

const Dashboard = (props) => {
    
    let selectOptions = () => (props.DataScience?[{value:"Media Mensal",dsmap:props.DataScience.months,type:"Line",desc:"Media mensal de pagamento"},{value:"Organizações Mais Caras",type:"",dsmap:props.DataScience.orgs,desc:"Media das organizações com a folha de pagamento mais alta"},{value:"Cargos Mais Bem Pagos",dsmap:props.DataScience.pos,type:"",desc:"Cargos com maior remuneração"},{value:"Faixa Salarial",dsmap:props.DataScience.hist,type:"Specific",desc:"Quantidade por faixa salarial"}]:[])

    const [select, setSelect] = useState("Selecione um grafico para exibir")

    const [expander, setExpander] = useState(false)

    const handleClick = val => {
        setSelect(val)
        if(expander){
            setExpander(!expander)
        }
    }

    const handleExpander = e => {
        setExpander(!expander)
        console.log(expander)
    }

    useEffect(()=>{
        if (props.DataScience === undefined){
            if (props.Loading === false || props.Loading === undefined) {
                props.GetDS()
            }
        }
    })

    const RenderMe = () => (<div className="Dash"><div className="options"><h1>Graficos</h1>
    <span onClick={handleExpander} className="selected">{select}</span>
    <SelecterList Expanded={expander} Selected={select} onPress={handleClick} Options={selectOptions()}/>
    </div>
    <ItemList Selected={select} Options= {selectOptions()}/>
    </div>)

    return (<Loading Loaded={RenderMe}/>)
}

const mapStateToProps = state => {
    return {
        DataScience: state.API.DataResum,
        Loading: state.API.Loading
    }
}

const mapDispatchToProps = dispatch => {
    return {
        GetDS: () => dispatch({type: REQUEST_FETCH, endpoint: "DataScience"})
    }
}

export default connect(mapStateToProps,mapDispatchToProps)(Dashboard);