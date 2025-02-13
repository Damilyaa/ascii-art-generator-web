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

function updateAsciiOutput() {
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
}

document.getElementById("text-input").addEventListener("input", updateAsciiOutput);

document.getElementById("banner-select").addEventListener("change", updateAsciiOutput);

document.addEventListener("DOMContentLoaded", function() {
    if (document.getElementById("text-input").value.trim()) {
        updateAsciiOutput();
    }
});

document.getElementById("banner-select").addEventListener("change", function() {
    const allowedBanners = ["standard", "shadow", "thinkertoy"];
    if (!allowedBanners.includes(this.value)) {
        alert("Invalid banner selected!");
        this.value = "standard"; // Reset to default
    }
});

document.getElementById("ascii-form").addEventListener("submit", function(event) {
    const text = document.getElementById("text-input").value.trim();
    const banner = document.getElementById("banner-select").value;

    const allowedBanners = ["standard", "shadow", "thinkertoy"];
    if (!allowedBanners.includes(banner)) {
        alert("Invalid banner selection!");
        event.preventDefault(); // Stop form submission
        return;
    }

    if (!text) {
        alert("Please enter some text!");
        event.preventDefault();
        return;
    }
});
