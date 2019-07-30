import React,{useState,useEffect,Fragment} from "react";
import PropTypes from "prop-types";
import {connect} from 'react-redux'
import '../static/notify.css'
import {CloseIcon} from './Icons'

const Notification = (props) => {
    
    let [type,setType] = useState("") 
    
    useEffect(()=>{
        if (type==="") {
            setType(props.type)
        }
    })

	const Classer = () => {
        let Class = "notifybox "
		switch (type) {
			case "Alert":
                Class += "alert"
				break;
			case "Info":
                Class += "info"
                break;
            case "None":
                Class += "none"
                break;
            default:
                Class += "basic"
                break;
        }
        return Class
    };
    const handleClick = e => {
        e.preventDefault();
        setType("None")
        props.handler()
    }
    return (<div onClick={handleClick} className={Classer()}><h3>{props.message}</h3>{CloseIcon()}</div>)
};

Notification.propTypes = {
    handler: PropTypes.func,
    message: PropTypes.string,
    type: PropTypes.string
}

const Notify = (props) => {
    return (<div className="general">{props.nfs.map((x)=><Notification key={x.uid} type={x.type} message={x.message} handler={() => props.removeNtf(x.uid)} />)}</div>);
};

const mapStateToProps = state => {
    return {
        nfs: state.Notifier
    }
} 

const mapDispatchToProps = dispatch => {
    return  {
        removeNtf: (id) => dispatch({type:"REMOVE_NOTIFY",index:id})
    }
}

export default connect(mapStateToProps,mapDispatchToProps)(Notify)