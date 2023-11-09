function updateAsciiArt() {
    let text = $('#text-input').val();

    // Check if the text is empty
    if (!text.trim()) {
        $('#output').text(''); // Clear the output
        return; // Exit the function
    }

    let banner = $("#banner-select").val();
    let windowWidth = window.innerWidth;

    $.ajax({
        url: "/ascii-art",
        type: "POST",
        data: { text: text, banner: banner, "window-width": windowWidth },
        success: function(result) {
            $('#output').text(result.AsciiArt);
        },
        dataType: 'json'
    });
}

document.getElementById('export-ascii-art').addEventListener('click', function() {
    var asciiArt = document.getElementById('output').textContent;
    if (!asciiArt.trim()) {
        alert("No content to download!");
        return;
    }
    var dataUri = 'data:text/plain;charset=utf-8,' + encodeURIComponent(asciiArt);
    var downloadLink = document.getElementById('download-link');
    downloadLink.href = dataUri;
    downloadLink.click();
});

