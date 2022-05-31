const options = {
	plotOptions: {
		series: {
			events: {
				legendItemClick: function (event) {
					event.preventDefault()
				}
			}
		}
	},
	chart: {
		backgroundColor: '',
		renderTo: 'graph',
		plotBorderColor: '#00ccff'
	},
	title: {
		text: ''
	},
	xAxis: {
		type: 'datetime',
		dateTimeLabelFormats: {
			day: '%a'
		}
	},
	yAxis: {
		title: {
			text: '',
			margin: 10
		}
	},
	credits: {
		enabled: false
	},
	series: [
		{
			type: 'area',
			name: 'Total messages',
			color: '#00ccff',
			data: []
		}
	]
}

const chart = new Highcharts.Chart(options)

function update_dstats(ttl_requests) {
	chart.series[0].addPoint([new Date().getTime(), ttl_requests], true, chart.series[0].points.length > 60)
}