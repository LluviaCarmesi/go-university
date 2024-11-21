export default function formatDateTime(dateTime: Date, forTable: boolean) {
    let dateTimeString = "";
    if (forTable) {
        return dateTimeString = dateTime.toLocaleDateString() + " " + dateTime.toLocaleTimeString();
    }
    dateTimeString += dateTime.getFullYear();
    dateTimeString += "-" + (dateTime.getMonth() < 10 ? "0" + dateTime.getMonth() + 1 : dateTime.getMonth() + 1);
    dateTimeString += "-" + (dateTime.getDate() < 10 ? "0" + dateTime.getDate() : dateTime.getDate());
    dateTimeString += "T" + (dateTime.getHours() < 10 ? "0" + dateTime.getHours() : dateTime.getHours());
    dateTimeString += ":" + (dateTime.getMinutes() < 10 ? "0" + dateTime.getMinutes() : dateTime.getMinutes());

    return dateTimeString;
}