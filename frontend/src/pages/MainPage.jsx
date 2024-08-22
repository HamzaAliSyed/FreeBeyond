import Card from "../components/Card/Card";
import { MainPageCards } from "../../frontenddata";

const MainPage = () => {
  return (
    <div className="p-4">
      {MainPageCards.map(({ id, title, image, description }) => (
        <Card
          key={id}
          title={title}
          description={description}
          image={image}
        />
      ))}
    </div>
  );
};

export default MainPage;
