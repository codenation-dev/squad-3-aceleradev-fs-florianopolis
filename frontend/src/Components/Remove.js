import React,{useState} from 'react'
import {REQUEST_FETCH} from '../redux/actions'
import {connect} from 'react-redux'
import '../static/remove.css'

const Remove = (props) => {

    let ID = parseInt(props.Fields.ID)

    let handleConfirm = e => {
        e.preventDefault()
        props.delete(ID)
        props.cancel()
    }

    return (<div className={props.Fields.RemoveModal?"SuperBox":"SuperBox none"}>
        <div className="remove">
        <h1>Você esta prestes a excluir {props.Fields.Name}. Tem certeza disso?</h1>
        <div id="TheButtons"><button id="Confirm" onClick={handleConfirm}>Excluir!</button><button id="Cancel" onClick={(e)=>{e.preventDefault(); props.cancel()}}>Cancelar!</button></div>
        </div>
    </div>)
}

const mapStateToProps = state => {
    return {Fields: state.User}
   }
   
   const mapDispatchToProps = dispatch => {
       return {
           delete: (ID) => dispatch({type:REQUEST_FETCH,endpoint:"Delete",args:{id:ID}} ),
           cancel: () => dispatch({type:"USER_MODAL",modal:""})
        }
   }
   

export default connect(mapStateToProps,mapDispatchToProps)(Remove)