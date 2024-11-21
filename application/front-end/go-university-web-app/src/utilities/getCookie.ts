export default function getCookie(property: string) {
    const name = property + "=";
    const decodedCookie = decodeURIComponent(document.cookie);
    const cookieProperties = decodedCookie.split(';');

    for (let i = 0; i < cookieProperties.length; i++) {
        let cookieProperty = cookieProperties[i];
        while (cookieProperty.charAt(0) == ' ') {
            cookieProperty = cookieProperty.substring(1);
        }
        if (cookieProperty.indexOf(name) == 0) {
            return cookieProperty.substring(name.length, cookieProperty.length);
        }
    }
    return "";
}