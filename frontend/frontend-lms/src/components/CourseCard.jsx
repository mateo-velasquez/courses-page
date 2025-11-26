import { Link } from 'react-router-dom';
import { useState, useEffect } from 'react';
import Card from './Card.jsx';
import Button from './Button.jsx';

const CourseCard = ({ course, onEnroll, showEnrollButton = true, isEnrolled = false }) => {
  const [showSuccess, setShowSuccess] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const handleEnrollClick = async () => {
    setIsLoading(true);
    await onEnroll(course);
    setShowSuccess(true);
    setIsLoading(false);
    // Hide success message after 3 seconds
    setTimeout(() => setShowSuccess(false), 3000);
  };
  const formatPrice = (price) => {
    return new Intl.NumberFormat('es-AR', {
      style: 'currency',
      currency: 'ARS',
    }).format(price);
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('es-AR');
  };

  const renderStars = (rating) => {
    const stars = [];
    const fullStars = Math.floor(rating);
    const hasHalfStar = rating % 1 !== 0;

    for (let i = 0; i < fullStars; i++) {
      stars.push('⭐');
    }
    if (hasHalfStar) {
      stars.push('⭐');
    }
    
    return stars.join('');
  };

  return (
    <Card className="h-full flex flex-col">
      <Card.Body className="flex-1">
        {course.image_url && (
          <div className="mb-4 -mx-4 -mt-4">
            <img
              src={`http://localhost:8090${course.image_url}`}
              alt={course.course_name}
              className="w-full h-40 object-cover rounded-t-lg"
              onError={(e) => {
                e.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjEyMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMjAwIiBoZWlnaHQ9IjEyMCIgZmlsbD0iI2VlZSIvPjx0ZXh0IHRleHQtYW5jaG9yPSJtaWRkbGUiIHg9IjEwMCIgeT0iNjAiIHN0eWxlPSJmaWxsOiNhYWE7Zm9udC1zaXplOjE2cHg7Zm9udC1mYW1pbHk6QXJpYWw7Ij5ObyBpbWFnZTwvdGV4dD48L3N2Zz4=';
              }}
            />
          </div>
        )}
        <div className="mb-4">
          <h3 className="text-xl font-semibold mb-2">
            <Link
              to={`/course/${course.id}`}
              className="text-primary hover:text-primary-hover transition-all"
            >
              {course.course_name}
            </Link>
          </h3>
          
          <p className="text-gray text-sm mb-3 line-clamp-3">
            {course.description}
          </p>

          <div className="flex items-center gap-4 text-sm text-gray mb-3">
            <span>Duración: {course.duration}</span>
            <span>Inicio: {formatDate(course.init_date)}</span>
          </div>

          {course.categories && course.categories.length > 0 && (
            <div className="flex flex-wrap gap-2 mb-3">
              {course.categories.map((category, index) => (
                <span
                  key={index}
                  className="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded-md"
                >
                  {category}
                </span>
              ))}
            </div>
          )}

          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <span className="text-sm">{renderStars(course.rating)}</span>
              <span className="text-sm text-gray">
                ({course.rating.toFixed(1)})
              </span>
            </div>
            <div className="text-lg font-bold text-primary">
              {formatPrice(course.price)}
            </div>
          </div>
        </div>
      </Card.Body>

      <Card.Footer>
        <div className="flex gap-2">
          <Link
            to={`/course/${course.id}`}
            className="btn btn-outline flex-1"
          >
            Ver Detalles
          </Link>
          {showEnrollButton && onEnroll && !isEnrolled && (
            <Button
              onClick={handleEnrollClick}
              variant="primary"
              className="flex-1"
              disabled={isLoading}
              loading={isLoading}
            >
              {isLoading ? 'Inscribiendo...' : 'Inscribirse'}
            </Button>
          )}
          {isEnrolled && (
            <button
              disabled
              className="px-4 py-2 bg-gray-300 text-gray-600 rounded-lg text-center font-medium flex-1 cursor-not-allowed opacity-60 border border-gray-400"
            >
              ✓ Ya estás inscripto
            </button>
          )}
        </div>
        {showSuccess && !isEnrolled && (
          <div className="mt-3 px-4 py-2 bg-green-100 text-green-700 rounded-lg text-center font-medium animate-pulse">
            ✓ Te has inscrito exitosamente
          </div>
        )}
      </Card.Footer>
    </Card>
  );
};

export default CourseCard;
