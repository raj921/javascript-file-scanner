document.getElementById('scanForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const url = document.getElementById('urlInput').value;
    const statusText = document.getElementById('statusText');
    const results = document.getElementById('results');
    const resultsContent = document.getElementById('resultsContent');

    statusText.textContent = 'Scanning...';
    results.classList.add('hidden');

    fetch('/scan', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `url=${encodeURIComponent(url)}`
    })
    .then(response => response.json())
    .then(data => {
        statusText.textContent = 'Completed';
        resultsContent.textContent = JSON.stringify(data, null, 2);
        results.classList.remove('hidden');
    })
    .catch(error => {
        statusText.textContent = 'Error';
        resultsContent.textContent = `An error occurred: ${error.message}`;
        results.classList.remove('hidden');
    });
});