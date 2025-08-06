export const truncateLength = (description) => {
    const maxLength = 10
    return description.length > maxLength
        ? description.slice(0, maxLength) + '...'
        : description
}