import moment from "moment";

export const noDataHandler = (value: string | number | null) : string | number => {
    return !value ? "нет данных" : value
}

export const parseDate = (value: string) : string => {
    return moment(value).format().replace("T", " ");
}