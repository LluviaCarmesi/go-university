export default function createChoices(items: any[], valueProperty: string, labelProperty: string) {
    let choices = [];

    for (let i = 0; i < items.length; i++) {
        choices.push({
            value: items[i][valueProperty],
            label: items[i][labelProperty]
        });
    }
    return choices;
}