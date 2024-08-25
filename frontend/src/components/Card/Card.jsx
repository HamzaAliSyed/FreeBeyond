const Card = ({ title, description, image, onClick }) => {
    return (
      <div 
        className="p-4 border rounded shadow hover:bg-gray-100 cursor-pointer"
        onClick={onClick}
      >
        <img src={image} alt={title} className="w-full h-auto" />
        <h2 className="text-lg font-bold">{title}</h2>
        <p>{description}</p>
      </div>
    );
  };
  
  export default Card;
  