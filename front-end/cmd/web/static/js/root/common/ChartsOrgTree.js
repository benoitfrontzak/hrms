class ChartsOrgTree {

    drawOrgTree(dataDT, target, size='medium'){
        
        const data = google.visualization.arrayToDataTable(dataDT), // Populate DT
              chart = new google.visualization.OrgChart(document.getElementById(target)), // Create the chart
              options = {
                'allowHtml':true,
                'allowCollapse':true,
                'nodeClass':'chartCell',
                'selectedNodeClass':'bg-secondary',
                'size':size // small | medium | large
              }

        // data.setRowProperty(3, 'style', 'border: 1px solid green')
        // data.setRowProperty(3, 'selectedStyle', 'background-color:#00FF00');
        
        // Draw the chart, setting the allowHtml option to true for the tooltips.
        chart.draw(data, options)
        
    }

}