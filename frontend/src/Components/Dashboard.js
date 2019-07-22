import React,{useState, Fragment} from 'react';
import '../static/dashboard.css';

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
    let itensList = props.Options.map(x=><Item selected={props.Selected===x.value} Render={<h2>example {x.value}</h2>}/>)
    return (<div className="FilterContainer">{itensList}</div>)
}    

const Item = (props) => {
    let show = () => (props.selected?"ActiveItem":"FilterItem")
    return(<div className={show()}>{props.Render}</div>)
}    

const Dashboard = (props) => {
    
    const [select, setSelect] = useState("None")

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

    return (<div className="Dash"><div className="options"><h1>Filter</h1>
    <span onClick={handleExpander} className="selected">{select}</span>
    <SelecterList Expanded={expander} Selected={select} onPress={handleClick} Options={[{value:"Test1"},{value:"Test2"},{value:"Test3"}]}/>
    </div>
    <ItemList Selected={select} Options={[{value:"Test1"},{value:"Test2"},{value:"Test3"}]}/>
    </div>);
}

export default Dashboard;