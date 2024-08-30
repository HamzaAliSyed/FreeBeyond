import { useNavigate } from "react-router-dom";

const Card = ({ image, name, description, link }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(link);
  };
  return (
    <div
      className="cursor-pointer rounded-lg shadow-md overflow-hidden"
      onClick={handleClick}
    >
      <img className="w-full h-72 object-cover" src={image} alt={name} />
      <div className="p-4">
        <h3 className="font-bold text-xl mb-2">{name}</h3>
        <p className="text-gray-700">{description}</p>
      </div>
    </div>
  );
};

export default Card;
