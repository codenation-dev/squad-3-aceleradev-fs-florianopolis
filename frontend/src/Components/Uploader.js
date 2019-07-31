import React,{useState} from 'react';
import {connect} from 'react-redux'
import '../static/upload.css'
import {CloseIcon,FileIcon,UpArrow} from '../Helpers/Icons';

const Uploader = (props) => {
   
    const [buffer,setBuffer] = useState("")
    const [fileName,setName] = useState("")
    
    const handleFile = e => {
        let files = e.target.files
        setName(files[0].name)
        let reader = new FileReader();
        reader.readAsDataURL(files[0])
        reader.onload=e=>{
            setBuffer(e.target.result)
        }
        
    }

    let SelectFile = () => {
        document.getElementById("FileSelecter").click()
    }

    return (
        <div className={props.Uploader?"SuperBox":"none"}>
        <div className="uploader">
            <div className="Head">
            <h3>Upload Client List</h3>
            <span onClick={props.closeUploader}>{CloseIcon()}</span>
            </div>
            <div className="Body">
            <p>{fileName?`Selected ${fileName}`:`Select CSV a List to upload!`}</p><span onClick={SelectFile}>{FileIcon()}{UpArrow()}</span>
            <input id="FileSelecter" onChange={handleFile} type="file" name="file" />
            </div>
            <button onClick={()=> {props.loadFile(buffer)}}>Upload Client List </button>
        </div>
        </div>
    )
}

const mapStateToProps = state => {
    return {
        Uploader: state.Uploader.Classer
    }
}

const mapDistpatchToProps = dispatch => {
    return {
        loadFile: (buffer) => dispatch({type:"UPLOAD_CSV_FILE",buffer:buffer}),
        closeUploader: () => dispatch({type:"CLOSE_UPLOADER"})
    }
}

export default connect(mapStateToProps,mapDistpatchToProps)(Uploader)