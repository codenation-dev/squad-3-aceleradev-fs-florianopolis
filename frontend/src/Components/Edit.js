import React,{useEffect,useState} from 'react'
import {connect} from 'react-redux'
import {REQUEST_FETCH} from '../redux/actions'
import '../static/edit.css'
import {CloseIcon} from '../Helpers/Icons'

const Edit = (props) => {

    let [ctr,SetCtr] = useState(0)
    let [name,setName] = useState("")
    let [mail,setMail] = useState("")
    let [temp,setTemp] = useState("")
    let [temp2,setTemp2] = useState("")
    let [pass,setPass] = useState("")
    let [required,setReq] = useState(false)

    let Classer = () => {
        if (props.Fields.EditModal===true){
            return "SuperBox"
        } else {
            return "none"
        }
    }

    let Closer = () => {
        setName("")
        setMail("")
        props.close()
        SetCtr(0)
    }

    const handleChange = e => {
        switch(e.target.id) {
            case "Name":
                setName(e.target.value)
                break;
            case "Email":
                setMail(e.target.value)
                break;
            case "Pass":
                setTemp(e.target.value)
                break;
            case "ConfirmPass":
                setTemp2(e.target.value)
                break;
            default:
                break;
        }
    }

    useEffect(()=>{
    
        if (ctr===0 && props.Fields.EditModal){
            setName(props.Fields.Name)
            setMail(props.Fields.Mail)
            SetCtr(1)
        }
        
        if (temp===temp2) {
            setPass(temp)
        } else {
            setPass("")
            let req = (temp.length != 0 || temp2.length != 0)?true:false
            setReq(req)
        }
    
    })

    let Fields = {id:parseInt(props.Fields.ID),usuario:name,email:mail,senha:pass}

    return (<div className={Classer()}><div className="UpdateForm">
        <div className="Options"><h1>User#{props.Fields.ID}</h1><span onClick={Closer}>{CloseIcon()}</span></div>
        <form>
            <label for="Name">Name</label>
            <input required type="text" id="Name" value={name} onChange={handleChange}/>
            <label for="Email">Email</label>
            <input required type="Email" id="Email" value={mail} onChange={handleChange}/>
            <label for="Pass" >Password</label>
            <input required={required} id="Pass" type="Password" value={temp} onChange={handleChange}/>
            <label for="ConfirmPass">Confirm Password</label>
            <input required={required} id="ConfirmPass" type="Password" value={temp2} onChange={handleChange}/>
            <button type="submit" onClick={e=> {e.preventDefault(); props.update(Fields);Closer() }}>Salvar Mudanças</button>
        </form>
            </div></div>)
}

const mapStateToProps = state => {
 return {Fields: state.User,
    Loading: state.API.Loading
    }
}

const mapDispatchToProps = dispatch => {
    return {
        update: (Fields) => dispatch({type:REQUEST_FETCH,endpoint:"Update",args:Fields}),
        close: () => dispatch({type:"USER_MODAL",modal:""})
    }
}

export default connect(mapStateToProps,mapDispatchToProps)(Edit)