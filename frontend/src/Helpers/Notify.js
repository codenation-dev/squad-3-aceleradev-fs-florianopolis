import React,{useState,useEffect,Fragment} from "react";
import PropTypes from "prop-types";
import {connect} from 'react-redux'
import '../static/notify.css'

const CloseIcon = () => (<svg id="i-close" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="32" height="32" fill="none" stroke="currentcolor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2">
<path d="M2 30 L30 2 M30 30 L2 2" />
</svg>)

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