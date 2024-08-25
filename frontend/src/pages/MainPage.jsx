import Card from "../components/Card/Card";
import { MainPageCards } from "../../frontenddata";
import { useNavigate } from "react-router-dom";

const MainPage = () => {
  const navigate = useNavigate();

  return (
    <div className="p-4 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      {MainPageCards.map(({ id, title, image, description, link }) => (
        <div key={id} className="col-span-1">
          <Card
            title={title}
            description={description}
            image={image}
            onClick={() => navigate(link)} 
          />
        </div>
      ))}
    </div>
  );
};

export default MainPage;
