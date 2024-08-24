import Card from "../components/Card/Card";
import { MainPageCards } from "../../frontenddata";
import { useNavigate } from "react-router-dom"

const MainPage = (
) => {
  const CardNavigate = useNavigate()
  return (
    <div className="p-4">
      {MainPageCards.map(({ id, title, image, description }) => (
        <a href='#'  onClick={(e) => {
          e.preventDefault()
          CardNavigate("/create-source")
        }}>
        <Card
          key={id}
          title={title}
          description={description}
          image={image}
        />
        </a>
      ))}
    </div>
  );
};

export default MainPage;
