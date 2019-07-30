import React from "react";
import Recharts, {
    BarChart,
    RadialBarChart,
    RadialBar,
    Pie,
    PieChart,
    FunnelChart,
    Funnel,
	AreaChart,
	Area,
	Tooltip,
	CartesianGrid,
	XAxis,
    YAxis,
    Legend,
    Bar,
    Line,
    LineChart,
    ResponsiveContainer,
    LabelList
} from "recharts";

let simpleData = (name,value) => {
    let nv = Math.round(value)
    return {name: name,value:nv}}

const Charter = (props) => {
    let Treated = Treat(props.data)
    let height = 400;
    switch (props.type) {
        case "Pie":
            return <TPieChart height={height} data={Treated}/>
        case "Radial":
            return <RadialType height={height} data={Treated} />
        case "Line":
            return <LineC height={height} data={Treated} />
        case "Specific":
            return <LineC height={height} data={Tclear(props.data)} />
        default:
            return <SimpleBar height={height} data={Treated}/>
    }
    
   
}

let Tclear = (data) => {
    let Arr = []
    for(var i in data){
        var key = i;
        let keyarr = key.split(" ")
        console.log(keyarr)
        console.log(parseFloat(keyarr[0]))
        let key1 = parseFloat(keyarr[0]).toFixed(2)
        let key2 = parseFloat(keyarr[2]).toFixed(2)
        
        let newkey = key1 + " to " + key2
        if (data[i]>100){
        Arr.push(simpleData(newkey,data[i]))
    }
    }
        return Arr
}


const SimpleBar = (props) => {
    
    return (
    <ResponsiveContainer width="90%" height={props.height}>
	<BarChart data={props.data}>
		<CartesianGrid strokeDasharray="3 3" />
		<XAxis dataKey="name" />
		<YAxis />
		<Tooltip />
		<Legend />
		<Bar dataKey="value" fill="#d68b00" >
        <LabelList fill="#fff" dataKey="name" position="center" angle="270" />
        </Bar>
	</BarChart>
    </ResponsiveContainer>
)}

const RadialType = props => {

    return (
        <ResponsiveContainer width="90%" height={props.height}>
    <RadialBarChart 
        innerRadius="10%" 
        outerRadius="80%" 
        data={props.data} 
        startAngle={180} 
        endAngle={0}
      >
        <RadialBar minAngle={15} label={{ fill: '#666', position: 'insideStart' }} background clockWise={true} dataKey='value' />
        <Legend iconSize={10} width={120} height={140} layout='vertical' verticalAlign='middle' align="right" />
        <Tooltip />
      </RadialBarChart></ResponsiveContainer>)
}

const TPieChart = (props) => {
    return (
        <ResponsiveContainer width="90%" height={props.height}>
    <PieChart >
     <Pie data={props.data} dataKey="value" nameKey="name" innerRadius={60} outerRadius={80} fill="#82ca9d" label/>
    </PieChart></ResponsiveContainer>
    )
}

const LineC = (props) => (<ResponsiveContainer width="90%" height={props.height}><LineChart data={props.data}
    margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
    <CartesianGrid strokeDasharray="3 3" />
    <XAxis dataKey="name" />
    <YAxis />
    <Tooltip />
    <Legend />
    <Line type="monotone" dataKey="value" stroke="#8884d8" />
  </LineChart></ResponsiveContainer>)

const Treat = (data) => {
    let Arr = []
    for (let key in data) {
        Arr.push(simpleData(key,data[key]))
    }
    return Arr
}

const FunelT = (props) => (
<ResponsiveContainer width="90%" height={props.height}>
    <FunnelChart >
  <Tooltip />
  <Funnel
    dataKey="value"
    data={props.data}
    isAnimationActive
  >
    <LabelList position="right" fill="#000" stroke="none" dataKey="name" />
  </Funnel>
</FunnelChart>
</ResponsiveContainer>
)

export default Charter;