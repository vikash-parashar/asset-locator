document.getElementById("getDataButton").addEventListener("click", () => {
    // Make an AJAX call to your Go backend API to get data
    fetch("/api/getdata")
        .then(response => response.json())
        .then(data => {
            populateTable("powerTable", data.powerDetails);
            populateTable("fiberTable", data.fiberDetails);
            populateTable("ownerTable", data.amcDetails);
            populateTable("locationTable", data.locationDetails);
        });
});

document.getElementById("addDataButton").addEventListener("click", () => {
    // Make an AJAX call to your Go backend API to add data
    fetch("/api/adddata")
        .then(response => response.json())
        .then(data => {
            // Handle the response or update the UI as needed
        });
});

function populateTable(tableId, data) {
    const table = document.getElementById(tableId);
    table.innerHTML = ""; // Clear the table content

    if (data.length === 0) {
        // Display a message when there's no data
        const row = table.insertRow();
        const cell = row.insertCell();
        cell.textContent = "No data available.";
    } else {
        // Populate the table with data
        // You can loop through the data and add rows and cells here
    }
}
