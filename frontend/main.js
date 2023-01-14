function fetchRentals(useFilter) {
    let filter = '';
    if(useFilter) {
        filter = '?customer=' + document.getElementById('FilterInput').value;
    }
    fetch('http://localhost:3000/rentals'+filter)
    .then((response) => response.json())
    .then((data) => {
        console.log(data)
        let htmlContent = '<table><tr><th>Created at</th><th>Updated at</th><th>Video</th><th>Customer</th><th>Status</th><th></th><th></th></tr>';
        
        data.rentals.forEach(rental => {
            const statusBgColor = rental.Status === 'available' ? '#95dd94' : '#f18383'; //green or red
            htmlContent += '<tr id='+rental.RentalID+'><td>'+rental.CreatedAt+'</td><td>'+rental.UpdatedAt+'</td><td>'+rental.VideoName+'</td><td>'+rental.Customer+'</td><td style="background-color:'+statusBgColor+'">'+rental.Status+'</td>'+
            '<td><button onclick="returnRental(\''+rental.RentalID+'\')">Return rental</button></td>'+
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

function returnRental(rentalID) {
    fetch('http://localhost:3000/rental/'+rentalID+'/return', { method: 'PUT'})
    .then((response) => response.json())
    .then((data) => {
        console.log(data);
        fetchRentals();
    });
}

function addRental() {
    const videoName = document.getElementById('VideoNameInput').value;
    const customer = document.getElementById('CustomerInput').value;

    fetch('http://localhost:3000/rental', { 
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body: JSON.stringify({ "VideoName": videoName, "Customer": customer })
    })
    .then((response) => response.json())
    .then((data) => {
        console.log(data);
        fetchRentals();
    }).catch((err) => {
        console.log(err);
        });;
}