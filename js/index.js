$(function() {
    alert('Hello from jQuery core function!');
    $(window).on('load', function() {
        alert('window on load body!')
    });
    $('#button01').on('click', function() {
        alert('Button clicked!');
    });
});