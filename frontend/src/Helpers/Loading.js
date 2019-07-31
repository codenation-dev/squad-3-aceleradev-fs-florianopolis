import React from 'react'
import '../static/loading.css'
import {connect} from 'react-redux';

const Loading = props => {
    const RenderIt = () => (<div class="lds-dual-ring"></div>)
    return (props.Loading?RenderIt():props.Loaded())
}

const mapStateToProps = state => {
    return {
        Loading: state.API.Loading
    }
}

export default connect(mapStateToProps)(Loading);