import CardTitle from "./CardTitle"
import CardDescription from "./CardDescription"
import CardCoverPhoto from "./CardCoverPhoto"
const Card = ({title, description,image}) => {
return (
    <div className="max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
        <div>
            {<CardCoverPhoto coverphoto={image} />}
            {<CardTitle title={title} />}
            {<CardDescription description={description} />}
        </div>
    </div>
)
}

export default Card