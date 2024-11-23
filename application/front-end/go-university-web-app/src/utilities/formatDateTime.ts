export default function formatDateTime(dateTime: Date, includeTime: boolean, forTable: boolean) {
    let dateTimeString = "";
    if (forTable) {
        return dateTimeString = dateTime.toLocaleDateString() + " " + (includeTime ? dateTime.toLocaleTimeString() : "");
    }
    dateTimeString += dateTime.getFullYear();
    dateTimeString += "-" + (dateTime.getMonth() < 10 ? "0" + (dateTime.getMonth() + 1) : dateTime.getMonth() + 1);
    dateTimeString += "-" + (includeTime ? dateTime.getDate() < 10 ?
        "0" + dateTime.getDate() :
        dateTime.getDate() :
        dateTime.getUTCDate() < 10 ?
            "0" + dateTime.getUTCDate() :
            dateTime.getUTCDate());
    if (includeTime) {
        dateTimeString += "T" + (dateTime.getHours() < 10 ? "0" + dateTime.getHours() : dateTime.getHours());
        dateTimeString += ":" + (dateTime.getMinutes() < 10 ? "0" + dateTime.getMinutes() : dateTime.getMinutes());
    }

    return dateTimeString;
}