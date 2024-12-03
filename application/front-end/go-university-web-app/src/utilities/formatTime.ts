export default function formatTime(time: string) {
    let formattedTime = "";
    const timeParts = time.split(":");
    const hour = parseInt(timeParts[0]);
    formattedTime = hour % 12 + ":" + timeParts[1] + " ";
    if (hour > 12) {
        formattedTime += "PM"
    }
    else {
        formattedTime += "AM";
    }

    return formattedTime;
}