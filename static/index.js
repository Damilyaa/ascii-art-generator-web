document.getElementById("text-input").addEventListener("input", function() {
    const text = document.getElementById("text-input").value;
    const banner = document.getElementById("banner-select").value;
    const output = document.getElementById("ascii-output");
    if (!text.trim()) {
        output.textContent = "Please enter text!";
        return;
    }
    fetch("/ascii-art", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: new URLSearchParams({ text, banner })
    })
    .then(response => response.text())
    .then(data => {
        output.textContent = data; 
    })
    .catch(error => {
        output.textContent = "Error: " + error;
    });
});