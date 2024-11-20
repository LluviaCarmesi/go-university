export default function getCookie(property: string) {
    const name = property + "=";
    const decodedCookie = decodeURIComponent(document.cookie);
    const cookieProperties = decodedCookie.split(';');

    for (let i = 0; i < cookieProperties.length; i++) {
        let c = cookieProperties[i];
        console.log(c);
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}