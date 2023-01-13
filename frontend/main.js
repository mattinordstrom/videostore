function fetchAllRentals() {
    fetch('http://localhost:3000/rentals')
    .then((response) => response.json())
    .then((data) => {
        console.log(data)
        let htmlContent = '<table><tr><th>Created at</th><th>Updated at</th><th>Video</th><th>Customer</th><th>Status</th><th></th><th></th></tr>';
        
        data.rentals.forEach(rental => {
            const statusBgColor = rental.Status === 'available' ? '#95dd94' : '#f18383'; //green or red
            htmlContent += '<tr id='+"41b6e958-8f02-4a59-8fa9-c32b841e2bba"+'><td>'+rental.CreatedAt+'</td><td>'+rental.UpdatedAt+'</td><td>'+rental.VideoName+'</td><td>'+rental.Customer+'</td><td style="background-color:'+statusBgColor+'">'+rental.Status+'</td>'+
            '<td><button>Return rental</button></td>'+
            '<td><button onclick="getPDF(\''+rental.RentalID+'\')">PDF</button></td></tr>';
        });
        
        htmlContent += '</table>';
        document.getElementsByClassName('content')[0].innerHTML = htmlContent;
    });
}

function getPDF(rentalID) {
    fetch('http://localhost:3000/rental/receipt/'+rentalID)
        .then((response) => response.blob())
        .then((blob) => {
        const _url = window.URL.createObjectURL(blob);
        window.open(_url, '_blank');
        }).catch((err) => {
        console.log(err);
        });
}