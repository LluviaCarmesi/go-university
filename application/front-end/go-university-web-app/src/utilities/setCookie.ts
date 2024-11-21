export default function setCookie(property: string, value: string, expiration: number) {
    const d = new Date();
    d.setTime(d.getTime() + (expiration * 24 * 60 * 60 * 1000));
    let expires = "expires=" + d.toUTCString();
    document.cookie = property + "=" + value + ";" + expires + ";path=/";
}